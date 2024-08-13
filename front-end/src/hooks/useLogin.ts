import { useContext, useState } from "react";
import axiosBase from "../config/axiosConfig";
import { AuthContext } from "../context/AuthContext";

interface AuthResponse {
  token: string,
  userId: string
}

const useLogin = () => {
  const { authDispatch } = useContext(AuthContext)
  const [logInError, setLogInError] = useState<String>("");

  const login = async (email: string, password: string) => {
    try {
      const response = await axiosBase.post<AuthResponse>(`/auth/login`, { email: email, password : password });
      // console.log(response.data.token)
      const user = { userID: response.data.userId, token: response.data.token, loggedIn: true}
      localStorage.setItem("auth_context_state", JSON.stringify(user));
      authDispatch({ type: 'LOG_IN', payload: user });
    } catch (err) {
        setLogInError("Invalid username or password")
    }
    // authContext.setEm
  }
  return {login, logInError}
}

export default useLogin