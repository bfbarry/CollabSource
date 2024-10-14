import { AxiosInstance } from 'axios';
import { Dispatch, SetStateAction, createContext, useEffect, useReducer } from 'react';
import axiosBase from '../config/axiosConfig';

// interface AuthContextI {
//     loggedIn: Boolean, 
//     setLoggedIn: Dispatch<SetStateAction<boolean>>;
//     // email: string,
//     // setEmail: Dispatch<SetStateAction<string>>;
// }

// export const AuthContext = createContext<AuthContextI>({
//     loggedIn: (localStorage.getItem("access_token") == null ? false : true), 
//     setLoggedIn:() => {},
//     // email: "",
//     // setEmail:() => {},
// });

interface ContextState {
    userID: string,
    token: string,
    loggedIn: boolean,
    // axiosBase: AxiosInstance,
}

// const init_state: ContextState = {userID: "public", token: "public", loggedIn: false, axiosBase: axiosBase}
const init_state: ContextState = {userID: "public", token: "public", loggedIn: false}
//useContext returns object with all ContextState values and authDispatch
export const AuthContext = createContext({...init_state, authDispatch: (() => {}) as Dispatch<any>});
interface ContextAction {
    type: string,
    payload: ContextState //TODO
}
const authReducer = (state: ContextState, action: ContextAction) =>  {
    switch (action.type) {
        case 'LOG_IN':
            axiosBase.defaults.headers.common['Authorization'] = action.payload.token
            axiosBase.defaults.headers.common['userId'] = action.payload.userID
            return { 
                ...action.payload
            }
        case 'LOG_OUT':
            axiosBase.defaults.headers.common['Authorization'] = "public"
            return init_state
        default:
            return state
    }
}

interface Props {
    children: React.ReactNode
}
export const AuthContextProvider = ({ children }: Props) => {
    const [state, authDispatch] = useReducer(authReducer, init_state)
    
    // only runs when this context renders
	useEffect(() => {
		const user = JSON.parse(JSON.stringify(init_state))
		// if (user.loggedIn ) {
            authDispatch({type: 'LOG_IN', payload: user})
        // }
	}, [])
    return (
        <AuthContext.Provider value={{...state, authDispatch}}>
            { children }
        </AuthContext.Provider>
    )
}
