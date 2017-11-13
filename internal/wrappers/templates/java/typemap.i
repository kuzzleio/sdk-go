%ignore all_statistics_result::res;

%typemap(javacode) all_statistics_result %{
  public Statistics[] getResult() {
    Statistics[] result = new Statistics[getRes_size()];
    for (int i = 0; i < result.length; ++i) {
      result[i] = getResult(i);
    }
    return result;
  }
%}

%javamethodmodifiers all_statistics_result::getResult(size_t pos) "private";
%extend all_statistics_result {
    statistics *getResult(size_t pos) {
        return $self->res + pos;
    }
}
