export interface ProjectBase {
  name       : string;
  description: string;
  category   : string;
  tags       : string[];
  seeking    : string[]; 
}

export interface FullProject extends ProjectBase {
  members    : member[];
  links      : string[];
  memberRequests: memberRequest[];
}

interface memberRequest {
  userId: string;
  name: string;
}

interface member {
  userId: string;
  name: string;
}

export interface ProjectWId extends ProjectBase {
  _id: string;
}

export interface Filters {
  categories: String[],
  searchQuery: String
}