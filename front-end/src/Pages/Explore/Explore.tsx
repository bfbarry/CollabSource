import React, { useEffect, useState } from "react";
import Select, { MultiValue } from 'react-select'
import axiosBase from "../../config/axiosConfig";
import { ProjectWId } from "../../types/project";
import ProjectTile from "../Common/ProjectTiles/ProjectTile";
import './Explore.css'
import { ReactComponent as RightSVG } from "../../assets/svg/right-next-navigation-svgrepo-com.svg"
import { ReactComponent as LeftSVG } from "../../assets/svg/left-navigation-back-svgrepo-com.svg"
import CreateProjectTile from "../Common/ProjectTiles/CreateProjectTile";


interface Filters {
  categories: String[],
  tags: String[]
}

interface OptionType {
  value: string;
  label: string;
}

const Explore: React.FC = () => {
  const [pageNum, setPageNum] = useState(1)
  const [projects, setProjects] = useState<ProjectWId[]>([])
  const categories = ['business', 'software engineering', 'art'] // TODO get from backend?
  const categorySelectOptions: OptionType[] = categories.map(opt => ({
    value: opt,
    label: opt
  }))
  const [filters, setFilters] = useState<Filters>({
    categories: [],
    tags: []
  })
  const NUMPERPAGE = 10

  const detectCategoryChange = (selected: MultiValue<OptionType>) => {
    const selectedValues = selected.map(o => o.value)
    let filterState = filters
    filterState.categories = selectedValues
    setFilters(filterState)
    setPageNum(1)
  }

  useEffect(() => {
    console.log('hello popsting', filters)
    axiosBase.post(`/projects?page=${pageNum}&size=${NUMPERPAGE}`, filters)
    .then(res => {
      setProjects(res.data.data)
    })
    .catch(err => {
      console.log(err)
    })
  }, [pageNum, filters])

  return (
    <>
      <div className='filter-parent'> 
        <h2>Filters</h2>
        <div className='filter-container'>
          <Select<OptionType, true>
            isMulti
            onChange={detectCategoryChange}
            options={categorySelectOptions}
            closeMenuOnSelect={false}
            className="basic-multi-select"
            classNamePrefix="select"
            />
        </div>
      </div>
      <div className='projectContainer'>
        <CreateProjectTile/>
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
      <div className='navigationContainer'>
        <div onClick={()=>setPageNum(pageNum-1)} className='navButton' id='prev-button'>
          {
            pageNum > 1 && 
            <>
              <LeftSVG className="SVGarrow"/>
              <button className="link-button"> previous </button>
            </>
          }
        </div>
        <div onClick={()=>setPageNum(pageNum+1)} className='navButton' id='next-button'>
          <button className="link-button"> next </button>
          <RightSVG className="SVGarrow"/>
        </div>
      </div>
    </>
  )
}

export default Explore

