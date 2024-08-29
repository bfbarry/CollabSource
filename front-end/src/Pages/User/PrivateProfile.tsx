import { FC } from "react";
import { UserType } from "../../types/user";

interface Props {
  user: UserType
}
const PrivateProfile:FC<Props> = ({ user }) => {
  return (
    <div>
      <h2> Profile Settings </h2>
      <hr/>
      <div> {user?.description} </div>
      {
        user?.skills &&
        <>
        <h2> Skills </h2>
        <hr/>
        <div> {user.skills.join(', ')} </div>
        </>
      }
  </div>
  )
}

export default PrivateProfile