import React, { useContext} from 'react'
import { AuthContext } from '../../context/AuthContext'
import {ProjectWithID} from "../../types/project"



const CreateProject:React.FC = () => {
  const { loggedIn } = useContext(AuthContext);

  return (
    <>
    { loggedIn? 
    <>hey baby {}</>:
    <p>who tf r u</p>
    }
    </>

  )
}

export default CreateProject