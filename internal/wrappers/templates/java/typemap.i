// Statistics[]
%ignore all_statistics_result::result;
%typemap(javacode) struct all_statistics_result %{
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

// CollectionsList[]
%ignore collection_entry_result::result;
%typemap(javacode) struct collection_entry_result %{
  public CollectionsList[] getResult() {
    CollectionsList[] result = new CollectionsList[(int)getResult_length()];
    for (int i = 0; i < result.length; ++i) {
      result[i] = getResult(i);
    }
    return result;
  }
%}

%javamethodmodifiers collection_entry_result::getResult(size_t pos) "private";
%extend collection_entry_result {
    collection_entry *getResult(size_t pos) {
        return $self->result + pos;
    }
}

// String[]
%ignore string_array_result::result;
%typemap(javacode) struct string_array_result %{
  public String[] getResult() {
    String[] result = new String[(int)getLength()];
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