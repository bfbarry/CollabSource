import React from 'react';
import './ProjectTile.css';
import { useNavigate } from "react-router-dom";
import { ProjectWId } from '../../../types/project';

const ProjectTile: React.FC<ProjectWId> = ({name, description, category, tags,_id, seeking}) => {

  let navigate = useNavigate();

  return(
    <div className="project-tile demo-project">
      <button 
      onClick={() => {navigate(`/project/${_id}`)}} 
      className='project-tile-button'>
        <div className='project-title'>
          <h3>{name}</h3>
        </div>
        <div className='description-container'>
          {description}
        </div>
        {/* <div className='tag-array'> */}
          <b>Seeking:</b> {seeking.join(', ')}
        {/* </div> */}
        {/* <div className='tag-array'> */}
          <b>Tags:</b> {tags.join(', ')}
        {/* </div> */}
      </button>
    </div>
  );
}

export default ProjectTile;