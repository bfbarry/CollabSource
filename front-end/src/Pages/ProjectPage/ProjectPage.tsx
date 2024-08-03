import { useEffect, useState } from "react";
import { useLocation, useParams } from "react-router-dom"
import axiosBase from "../../config/axiosConfig";
import "./ProjectPage.css"
import { ProjectBase } from "../../types/project";


const ProjectPage = () => {
  const { id } = useParams();
  const location = useLocation();
  const placeHolderProject = {
      _id: "",
      name:        "",
      description: "",
      category:    "",
      tags:        [],
      seeking:     [],
      ownerId:     ""
  }
  const [projectData, setProjectData] = useState<ProjectBase>(location.state || placeHolderProject);
  useEffect(() => {
    const dataFromNavigate = location.state != null
    if (!dataFromNavigate) {
      console.log('hello')
      axiosBase.get(`/project/${id}`)
      .then(res => {setProjectData(res.data.project)})
      .catch(err => {console.log(err)}) // TODO set UI error
      return
    }
  }, [id])

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
              ) )
            }
          </div>
        </div>
      </div>
    </div>
  )
}

export default ProjectPage