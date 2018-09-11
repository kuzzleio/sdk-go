%{
  #include <functional>
%}

#define param(num,type) $typemap(jstype,type) arg ## num
#define unpack(num,type) arg##num
#define lvalref(num,type) type&& arg##num
#define forward(num,type) std::forward<type>(arg##num)

// Iterate over MACRO Variadic
// https://stackoverflow.com/questions/1872220/is-it-possible-to-iterate-over-arguments-in-variadic-macros/11994395#11994395
#define FE_0(...)
#define FE_1(action,a1) action(0,a1)
#define FE_2(action,a1,a2) action(0,a1), action(1,a2)
#define FE_3(action,a1,a2,a3) action(0,a1), action(1,a2), action(2,a3)
#define FE_4(action,a1,a2,a3,a4) action(0,a1), action(1,a2), action(2,a3), action(3,a4)
#define FE_5(action,a1,a2,a3,a4,a5) action(0,a1), action(1,a2), action(2,a3), action(3,a4), action(4,a5)

#define GET_MACRO(_1,_2,_3,_4,_5,NAME,...) NAME
%define FOR_EACH(action,...) 
  GET_MACRO(__VA_ARGS__, FE_5, FE_4, FE_3, FE_2, FE_1, FE_0)(action,__VA_ARGS__)
%enddef

%define %std_function(Name, Ret, AbstractMethodName, ...)

%feature("director") Name##Impl;
%typemap(javaclassmodifiers) Name##Impl "abstract class";

%{
  struct Name##Impl {
    virtual ~Name##Impl() {}
    virtual Ret AbstractMethodName(__VA_ARGS__) = 0;
  };
%}

%javamethodmodifiers Name##Impl::AbstractMethodName "abstract protected";
%typemap(javaout) Ret Name##Impl::AbstractMethodName ";"

%typemap(javaclassmodifiers) std::function<Ret(__VA_ARGS__)> "public abstract class"
%javamethodmodifiers std::function<Ret(__VA_ARGS__)>::operator() "public abstract"
%typemap(javaout) Ret std::function<Ret(__VA_ARGS__)>::operator() ";"

struct Name##Impl {
  virtual ~Name##Impl();
protected:
  virtual Ret AbstractMethodName(__VA_ARGS__) = 0;
};

%typemap(maybereturn) SWIGTYPE "return ";
%typemap(maybereturn) void "";

%typemap(javain) std::function<Ret(__VA_ARGS__)>, std::function<Ret(__VA_ARGS__)>* "$javaclassname.getCPtr($javaclassname.makeNative($javainput))"
%typemap(javacode) std::function<Ret(__VA_ARGS__)> %{
  public Name() {
    wrapper = new Name##Impl(){
      public $typemap(jstype, Ret) AbstractMethodName(FOR_EACH(param, __VA_ARGS__)) {
        $typemap(maybereturn, Ret)Name.this.AbstractMethodName(FOR_EACH(unpack, __VA_ARGS__));
      }
    };
    proxy = new $javaclassname(wrapper){
      public $typemap(jstype, Ret) AbstractMethodName(FOR_EACH(param, __VA_ARGS__)) {
        $typemap(maybereturn, Ret)Name.this.AbstractMethodName(FOR_EACH(unpack, __VA_ARGS__));
      }
    };
  }

  static $javaclassname makeNative($javaclassname in) {
    if (null == in.wrapper) return in;
    return in.proxy;
  }

  private Name##Impl wrapper;
  private $javaclassname proxy;
%}

%rename(Name) std::function<Ret(__VA_ARGS__)>;
%rename(AbstractMethodName) std::function<Ret(__VA_ARGS__)>::operator();

namespace std {
  struct function<Ret(__VA_ARGS__)> {
    function<Ret(__VA_ARGS__)>(const std::function<Ret(__VA_ARGS__)>&);

    Ret operator()(__VA_ARGS__) const;

    function<Ret(__VA_ARGS__)>(Ret(*const)(__VA_ARGS__));

    %extend {
      function<Ret(__VA_ARGS__)>(Name##Impl *in) {
        return new std::function<Ret(__VA_ARGS__)>([=](FOR_EACH(lvalref,__VA_ARGS__)){
              return in->AbstractMethodName(FOR_EACH(forward,__VA_ARGS__));
        });
      }
    }
  };
}

%enddef