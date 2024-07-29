export interface ProjectBase {
  name       : string;
  description: string;
  category   : string;
  tags       : string[];
}

// TODO probably will rename once use cases expand
export interface ProjectWithID extends ProjectBase {
  id: string;
}

export interface ProjectForm extends ProjectBase {
  ownerEmail: string;
}
