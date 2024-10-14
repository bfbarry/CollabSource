import { useContext, useState } from "react";
import axiosBase from "../config/axiosConfig";
import { AuthContext } from "../context/AuthContext";

export interface AuthResponse {
  token: string,
  userId: string
}

const useLogin = () => {
  const { authDispatch } = useContext(AuthContext)
  const [logInError, setLogInError] = useState<String>("");

  const login = async (email: string, password: string) => {
    try {
      const response = await axiosBase.post<AuthResponse>(`/auth/login`, { email: email, password : password });

      const user = { userID: response.data.userId, token: response.data.token, loggedIn: true}
      // axiosBase.defaults.headers.common['Authorization']  = response.data.token
      // localStorage.setItem("auth_context_state", JSON.stringify(user));
      console.log("about to log in")
      authDispatch({ type: 'LOG_IN', payload: user });
    } catch (err) {
        setLogInError("Invalid username or password")
    }
    // authContext.setEm
  }
  return {login, logInError}
}

export default useLogin