import { FC, useContext, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { AuthContext } from "../../context/AuthContext";
import axiosBase from "../../config/axiosConfig";
import { AxiosHeaders } from 'axios';
import {  UserType  } from "../../types/user";
import './User.css'
import { ReactComponent as ProfileSVG } from "../../assets/svg/user-profile-filled-svgrepo-com.svg"

const User:FC = () => {
  const { id } = useParams();
  const { loggedIn, token } = useContext(AuthContext)
  const [ user, setUser] = useState<UserType>(null)
  
  useEffect(() => {
    let headers = new AxiosHeaders()
    if (loggedIn) {
      headers.Authorization = token
    } else {
      headers.Authorization = "public"
    }
    axiosBase.get(`/user/${id}`, { headers })
    .then(res => {setUser(res.data.data)})
    .catch(err => {console.log(err)}) // TODO set UI error
    return
  }, [id, loggedIn])
  return (
      
    <div>
      <div className='name-header'>
        <div className="svg-cont">
          <ProfileSVG className='prof-pic'/>
        </div>
            <h1> {user?.name} </h1>
      </div>
      <div className="profileBorder">
        <div className="profileBody">
          <h2> About </h2>
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
          <h2>Projects</h2>
          <hr/>
        </div>
      </div>
    </div>
  )
}

export default User