import React from 'react';
import './ProjectTile.css';
import { useNavigate } from "react-router-dom";
import { ProjectWId } from '../../../types/project';

const ProjectTile: React.FC<ProjectWId> = ({name, description, category, tags,_id, seeking}) => {

    let navigate = useNavigate();

    return(
        <div className="project-tile demo-project">
            <button onClick={() => {navigate(`/project/${_id}`)}} className='project-tile-button'>
                <h3>{name}</h3>
                <p className="project-tile-description">{description}</p>
                <p className="project-tile-description"><b>Seeking:</b> {seeking.join(', ')}</p>
            </button>
        </div>
    );
}

export default ProjectTile;