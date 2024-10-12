import { useContext, useEffect, useState } from "react";
import { Link, useLocation, useParams } from "react-router-dom"
import axiosBase from "../../config/axiosConfig";
import "./ProjectPage.css"
import { FullProject } from "../../types/project";
import { AuthContext } from "../../context/AuthContext";
import SignUpModal from "../Modals/SignUpModal";
import PromptAccount from "../Modals/PromptAccount";
import LogInModal from "../Modals/LogInModal";

interface ProjectMetaData {
  isPublic: boolean,
  isMember: boolean,
  isOwner: boolean,
  isPending: boolean
}


const ProjectPage = () => {
  const { userID, loggedIn } = useContext(AuthContext)
  const { id } = useParams();
  const location = useLocation();
  const placeHolderProject = {
      _id: "",
      name:        "",
      description: "",
      category:    "",
      tags:        [],
      seeking:     [],
      ownerId:     "",
      memberRequests: [],
      members: [],
  }

  const [projectData, setProjectData] = useState<FullProject>(location.state || placeHolderProject);
  const [projectMetaData, setProjectMetaData] = useState<ProjectMetaData>({isPublic: true, isMember: false, isOwner: false, isPending: false});
  const [showPromptAccount, setShowPromptAccount ] = useState<Boolean>(false)
  const [showLogin, setShowLogin ] = useState<Boolean>(false)
  const [showSignUp, setShowSignUp ] = useState<Boolean>(false)
  const [refreshPage, setRefreshPage] = useState<Boolean>(false);

  // useEffect(() => {
  //   const dataFromNavigate = location.state != null
  //   if (!dataFromNavigate) {
  //     axiosBase.get(`/project/${id}`)
  //     .then(res => {
  //       setProjectData(res.data.project)
  //       setProjectMetaData({isPublic: res.data.isPublic, isMember: res.data.isMember, isOwner: res.data.isOwner, isPending: res.data.isPending});
  //     })
  //     .catch(err => {console.log(err)}) // TODO set UI error
  //     return
  //   }
  // }, [id, location.state ])

  useEffect(() => {
    axiosBase.get(`/project/${id}`)
      .then(res => {
        setProjectData(res.data.project)
        setProjectMetaData({isPublic: res.data.isPublic, isMember: res.data.isMember, isOwner: res.data.isOwner, isPending: res.data.isPending});
      })
      .catch(err => {console.log(err)}) // TODO set UI error
      return
  },[loggedIn, id, refreshPage])

  const joinProject = () => {
    if(!loggedIn){
      setShowPromptAccount(true)
      return
    }
    axiosBase.post(`project/project_request/dummy`,{userId: userID , projectId: id})
      .then(res => {
        setProjectMetaData({isPublic: false, isMember: false, isOwner: false, isPending: true});
      })
      .catch(err => {console.log(err)})
  }

  const acceptMemberRequest = (action:string, userId: string, name: string) => {
    if(action === "Accept"){
      axiosBase.patch(`project/project_request/${id}`,{userId: userId , name: name, admission: "accepted"})
      .then(res => {
        setRefreshPage(!refreshPage)
      })
      .catch(err => {console.log(err)})
    } else {
      axiosBase.patch(`project/project_request/${id}`,{userId: userId , name: name , admission: "denied"})
      .then(res => {
        setRefreshPage(!refreshPage)
      })
      .catch(err => {console.log(err)})
    }
  }

  const renderLinksAndMembers = () => {

    if(projectData.links == null){
      projectData.links = []
    } 

    if(projectData.members == null){
      projectData.members = []
    } 

    if(projectData.memberRequests == null){
      projectData.memberRequests = []
    }
    
    return (
      <div>
        <div>
          <b> Links: </b>
          {
            projectData.links.map((v) => (
              <div>- {v} </div>
            ))
          }
        </div>
        <div>
          <b> Members: </b>
          <div className="ProjectMembers">
            {
              projectData.members.map((v) => (
                <Link to ={`/user/${v.userId}`}>
                  {v.name}
                </Link>
              ))
            }
          </div>
        </div>
        <div>
          {
          projectMetaData.isOwner ?
          <div>
          <b> Members Requests: </b>
          {
          projectData.memberRequests.map((v) => (
              <div id = "MembershipRequests" >
                <Link to ={`/user/${v.userId}`}>
                  {v.name}
                </Link>
                <button className="MembershipButton" onClick={() => acceptMemberRequest("Accept", v.userId, v.name)}>Accept</button>
                <button className="MembershipButton" onClick={() => acceptMemberRequest("Reject", v.userId, v.name)}>Reject</button>
              </div>
            ))
          }
          </div>
          : 
          null 
          }
        </div>
      </div>
    );
  }

  return (
    <div className='container'>
      <div className='header'>
        <h2> {projectData.name} </h2>
        <h3> a {projectData.category} project </h3>
      </div>
      <div className='projectBorder'>
        <div className='projectBody'>
          <div className='tagContainer'>
            <b>Tags: </b>
            {
              projectData.tags.map((v,i) => (
                <div key={i} className='tag'>{v}</div>
              ))
            }
          </div>
          <div>
            <p>{projectData.description}</p>
            <b>Seeking: </b>
            {
              projectData.seeking.map((v, i) => (
                <div>- {v} </div>
              ))
            }
          </div>
          {loggedIn && (projectMetaData.isMember || projectMetaData.isOwner) ? 
            renderLinksAndMembers()
          :
          projectMetaData.isPending ? 
            <div>
              <div id="requestButtonContainer">
                <button id="requestButton" disabled={true}>Request Pending</button>
              </div>
            </div>
            : 
            <div>
              <div id="requestButtonContainer">
                <button id="requestButton" onClick={joinProject}>Request to join</button>
              </div>
            </div>
          }
        </div>
      </div>
      {showPromptAccount &&
        <PromptAccount 
          setShowPromptAccount={setShowPromptAccount}
          setShowLogIn={setShowLogin}
          setShowSignUp={setShowSignUp}
          />
      }
      { showSignUp && <SignUpModal setShowSignUp={setShowSignUp}/> }
      { showLogin && <LogInModal SetShowLogIn={setShowLogin}/> }
    </div>
  )
}

export default ProjectPage