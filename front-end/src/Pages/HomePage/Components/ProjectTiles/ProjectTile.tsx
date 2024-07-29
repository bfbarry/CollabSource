import React from 'react';
import './ProjectTile.css';
import { useNavigate } from "react-router-dom";
import { ProjectBase } from "../../../../types/project"

interface ProjectTileProps extends ProjectBase {
  id: string;
}
const ProjectTile: React.FC<ProjectTileProps> = ({name, description, category, tags,id}) => {

    let navigate = useNavigate();

    return(
        <div className="project-tile demo-project">
            <button onClick={() => {navigate(`/project/${id}`)}} className='project-tile-button'>
                <h3>{name}</h3>
                <p className="project-tile-description">{description}</p>
                <p className="project-tile-description"><b>Seeking:</b> Position1, Position2</p>
            </button>
        </div>
    );
}

export default ProjectTile;