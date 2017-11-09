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

%typemap(javacode) collections_list_result %{
  public CollectionsList[] getResult() {
    CollectionsList[] result = new CollectionsList[getCollection_list_length()];
    for (int i = 0; i < result.length; ++i) {
      result[i] = getResult(i);
    }
    return result;
  }
%}

%javamethodmodifiers collections_list_result::getResult(size_t pos) "private";
%extend collections_list_result {
    collections_list *getResult(size_t pos) {
        return $self->res + pos;
    }
}