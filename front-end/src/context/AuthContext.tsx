import { Dispatch, SetStateAction, createContext, useEffect, useReducer } from 'react';

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
}

const init_state: ContextState = {userID: "", token: "", loggedIn: false}
//useContext returns object with all ContextState values and authDispatch
export const AuthContext = createContext({...init_state, authDispatch: (() => {}) as Dispatch<any>});

interface ContextAction {
    type: string,
    payload: ContextState //TODO
}
const authReducer = (state: ContextState, action: ContextAction) => {
    switch (action.type) {
        case 'LOG_IN':
            return { 
                ...action.payload
            }
        case 'LOG_OUT':
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
		const user = JSON.parse(localStorage.getItem("auth_context_state") || JSON.stringify(init_state))
		if (user.loggedIn ) {
            authDispatch({type: 'LOG_IN', payload: user})
        }
	}, [])
    return (
        <AuthContext.Provider value={{...state, authDispatch}}>
            { children }
        </AuthContext.Provider>
    )
}
