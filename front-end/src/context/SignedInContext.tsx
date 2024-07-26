import { Dispatch, SetStateAction, createContext } from 'react';

interface signedInContextType {
    signedInUser: Boolean, 
    setSignedInUser: Dispatch<SetStateAction<boolean>>;
}

export const SignedInContext = createContext<signedInContextType>({signedInUser: (localStorage.getItem("access_token") == null ? false : true), setSignedInUser:() => {}});