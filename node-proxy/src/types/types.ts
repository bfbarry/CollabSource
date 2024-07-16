export interface UserRegisterRequestBody {
    name: string;
    password: string;
    email: string;
    description: string;
    skills: string[];
}

export interface UserPatchRequestBody {
    name: string;
    description: string;
    skills: string[];
}

export interface Project {
    name       : string;
    description: string;
    category   : string;
    tags       : string[];
  }