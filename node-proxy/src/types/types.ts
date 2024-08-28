import { Request } from "express";

export interface UserRegisterRequestBody {
    name: string;
    password: string;
    email: string;
    description: string;
    skills: string[];
}

export interface UserPatchRequestBody {
    name : string;
    description : string;
    skills : string[];
}

export interface Project {
    id : string
    ownerId: string
    name : string
    description : string
    category : string
    tags : string[]
    seeking: string[]
  }


export interface PaginatedResponseBody<T> {
	items: T[]
	page: number
	hasNext: boolean
}

export interface IRequest extends Request {
    id: string
  }