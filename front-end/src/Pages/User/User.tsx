import { FC, useContext, useEffect } from "react";
import { useParams } from "react-router-dom";
import { AuthContext } from "../../context/AuthContext";

const User:FC = () => {
  const { id } = useParams();
  const { loggedIn, userID } = useContext(AuthContext)

  useEffect(() => {

  })
  return (
    {
      
    }
    <div>

    </div>
  )
}

export default User