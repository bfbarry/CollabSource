import { FC, useContext, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { AuthContext } from "../../context/AuthContext";
import axiosBase from "../../config/axiosConfig";
import { AxiosHeaders } from 'axios';
import {  UserType  } from "../../types/user";
import {  Filters, ProjectWId  } from "../../types/project";
import './User.css'
import { ReactComponent as ProfileSVG } from "../../assets/svg/user-profile-filled-svgrepo-com.svg"
import { ReactComponent as RightSVG } from "../../assets/svg/right-next-navigation-svgrepo-com.svg"
import { ReactComponent as LeftSVG } from "../../assets/svg/left-navigation-back-svgrepo-com.svg"
import ProjectTile from "../Common/ProjectTiles/ProjectTile";
import PublicProfile from "./PublicProfile";
import PrivateProfile from "./PrivateProfile";


const User:FC = () => {
  const { id } = useParams();
  const { loggedIn, userID, token } = useContext(AuthContext)
  const [ user, setUser] = useState<UserType>(null)
  const [projects, setProjects] = useState<ProjectWId[]>([]);
  const [pageNum, setPageNum] = useState(1)
  const [hasNext, setHasHext] = useState<Boolean>(false)
  
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
  
  useEffect(() => {
    axiosBase.get(`/user/projects/${id}?page=${pageNum}&size=4`)
    .then(res => {
      setProjects(res.data.items || [])
      setHasHext(res.data.hasNext)
    })
    .catch(error => {
      console.log(error)
    })
    return
  }, [id, pageNum])

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
          {
            loggedIn && userID === id ?
            <PrivateProfile user={user}/> :
            <PublicProfile user={user}/>

          }
          <h2>Projects</h2>
          <hr/>
          <div id="explore-projects-section">
            <div id="project-tiles-section">
            {projects.map((value) => (
              <ProjectTile 
              key={value._id}
              _id={value._id}
              name={value.name} 
              description={value.description} 
              category={value.category} 
              tags={value.tags} 
              seeking={value.seeking}
              />
              ))}
            </div>
          </div>
          <div className='navigationContainer'>
          {
            pageNum > 1 ? 
            <div onClick={()=>setPageNum(pageNum-1)} 
            className='navButton' id='prev-button'>
                <>
                  <LeftSVG className="SVGarrow"/>
                </> 
            </div>:
            <div className='arrow-placeholder'></div>
          }
          {
            hasNext ?
            <div onClick={()=>setPageNum(pageNum+1)} 
            className='navButton' id='next-button'>
                <>
                  <RightSVG className="SVGarrow"/>
                </>
            </div> :
            <div className='arrow-placeholder'></div>
          }
          </div>
        </div>
      </div>
    </div>
  )
}

export default User