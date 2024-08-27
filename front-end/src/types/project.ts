export interface ProjectBase {
  name       : string;
  description: string;
  category   : string;
  tags       : string[];
  seeking    : string[] 
}

export interface ProjectWId extends ProjectBase {
  _id: string;
}

export interface Filters {
  categories: String[],
  searchQuery: String
}