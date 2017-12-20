%include <typemap.i>

// Statistics[]
%ignore all_statistics_result::result;
%typemap(javacode) all_statistics_result %{
  public Statistics[] getResult() {
    Statistics[] result = new Statistics[(int)getResult_length()];
    for (int i = 0; i < result.length; ++i) {
      result[i] = getResult(i);
    }
    return result;
  }
%}

%javamethodmodifiers all_statistics_result::getResult(size_t pos) "private";
%extend all_statistics_result {
    statistics *getResult(size_t pos) {
        return $self->result + pos;
    }
}

%javamethodmodifiers collection_entry_result::getResult(size_t pos) "private";
%extend collection_entry_result {
    collection_entry *getResult(size_t pos) {
        return $self->result + pos;
    }
}

// String[]
%ignore string_array_result::result;
%typemap(javacode) string_array_result %{
  public String[] getResult() {
    String[] result = new String[(int)getResult_length()];
    for (int i = 0; i < result.length; ++i) {
      result[i] = getResult(i);
    }
    return result;
  }
%}

%javamethodmodifiers string_array_result::getResult(size_t pos) "private";
%extend string_array_result {
    char *getResult(size_t pos) {
        return *$self->result + pos;
    }
}

// Date
%ignore date_result::result;
%typemap(javacode) struct date_result %{
  public java.util.Date getResult() {
    return new java.util.Date(getDateResult());
  }
%}

%javamethodmodifiers date_result::getDateResult() "private";
%extend date_result {
    long long getDateResult() {
        return $self->result;
    }
}
