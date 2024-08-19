import React, { useEffect, useState, KeyboardEvent, useRef } from "react";
import Select, { MultiValue } from 'react-select'
import axiosBase from "../../config/axiosConfig";
import { ProjectWId } from "../../types/project";
import ProjectTile from "../Common/ProjectTiles/ProjectTile";
import './Explore.css'
import { ReactComponent as RightSVG } from "../../assets/svg/right-next-navigation-svgrepo-com.svg"
import { ReactComponent as LeftSVG } from "../../assets/svg/left-navigation-back-svgrepo-com.svg"
import CreateProjectTile from "../Common/ProjectTiles/CreateProjectTile";
import { Filters } from "../../types/project";
import SearchBar from "../Common/SearchBar";
import { useLocation } from "react-router-dom";

interface OptionType {
  value: string;
  label: string;
}
export const NUMPERPAGE = 10

const Explore: React.FC = () => {
  
  const location = useLocation()
  let [redirected, setRedirected] = useState(location.state != null)
  let locationProjects = null
  let locationHasNext = null
  let defaultSearchQuery = ''
  if (redirected) {
    locationProjects = location.state.projects
    locationHasNext = location.state.hasNext
    defaultSearchQuery = location.state.searchQuery
    console.log(defaultSearchQuery)
  }
  const [pageNum, setPageNum] = useState(1)
  const [projects, setProjects] = useState<ProjectWId[]>(locationProjects || [])
  const [hasNext, setHasHext] = useState<Boolean>(locationHasNext || false)
  const categoryChoices = ['business', 'software engineering', 'art'] // TODO get from backend?
  const categorySelectOptions: OptionType[] = categoryChoices.map(opt => ({
    value: opt,
    label: opt
  }))
  const [categories, setCategories] = useState<String[]>([])
  // TODO get state from front page search
  const [searchQuery, setSearchQuery] = useState<string>(defaultSearchQuery)
  
  const detectCategoryChange = (selected: MultiValue<OptionType>) => {
    const selectedValues = selected.map(o => o.value)
    setCategories(selectedValues);
    setPageNum(1)
  }

  useEffect(() => {
    if (!redirected) {
      const filters: Filters = { categories, searchQuery }
      axiosBase.post(`/projects?page=${pageNum}&size=${NUMPERPAGE}`, filters)
      .then(res => {
        setProjects(res.data.items || [])
        setHasHext(res.data.hasNext)
      })
      .catch(err => {
        console.log(err)
      })
    }
    setRedirected(false)
  }, [pageNum, categories])

  return (
    <>
      <div className='filter-parent'>
        <SearchBar 
          searchQuery={searchQuery} 
          setSearchQuery={setSearchQuery}
          categories={categories}
          setProjects={setProjects}
          setHasHext={setHasHext}
          redirect={false}/>
        <div className='filter-dropdowns'> 
          <h2>Filters</h2>
          <div className='filter-container'>
            <Select<OptionType, true>
              placeholder="category"
              isMulti
              onChange={detectCategoryChange}
              options={categorySelectOptions}
              closeMenuOnSelect={false}
              className="basic-multi-select"
              classNamePrefix="select"
              />
          </div>
        </div>
      </div>
      <div className='projectContainer'>
        <CreateProjectTile/>
        {projects.length > 0 ? projects.map((value) => (
          <ProjectTile 
          key={value._id}
          _id={value._id}
          name={value.name} 
          description={value.description} 
          category={value.category} 
          tags={value.tags} 
          seeking={value.seeking}
          />
      )) : 
      <div> no projects found through those filters (yet) </div>}
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
          {
            hasNext &&
            <>
              <button className="link-button"> next </button>
              <RightSVG className="SVGarrow"/>
            </>
          }
        </div>
      </div>
    </>
  )
}

export default Explore

