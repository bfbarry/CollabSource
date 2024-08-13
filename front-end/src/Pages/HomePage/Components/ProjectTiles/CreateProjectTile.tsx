import React from 'react';
import './ProjectTile.css';
import { useNavigate } from "react-router-dom";
import { ReactComponent as PlusSVG } from "../../../../assets/svg/plus-large-svgrepo-com.svg"

const CreateProjectTile: React.FC = () => {

    let navigate = useNavigate();

    return(
      <div className="project-tile create-project">
        <button onClick={() => {navigate(`/create_project`)}} className='project-tile-button'>
          <h3>Create a project</h3>
          <div className='plus-div'>

            <PlusSVG className="plus-sign"/>
          </div>
        </button>
      </div>
    );
}

export default CreateProjectTile;