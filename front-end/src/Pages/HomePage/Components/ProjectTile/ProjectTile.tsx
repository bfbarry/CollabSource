import React from 'react';
import './ProjectTile.css';
import { useNavigate } from "react-router-dom";

interface Props {
    id: string;
    name       : string;
    description: string;
    category   : string;
    tags       : string[];
}

const ProjectTile: React.FC<Props> = ({name, description, category, tags,id}) => {

    let navigate = useNavigate();

    return(
        <button onClick={() => {navigate(`/project/${id}`)}} id = 'project-tile-button'>
            <div id = "project-tile">
                <h3>{name}</h3>
                <p id="project-tile-description">{description}</p>
                <p id="project-tile-description"><b>Seeking:</b> Position1, Position2</p>
            </div>
         </button>
    );
}

export default ProjectTile;