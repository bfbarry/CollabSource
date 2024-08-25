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
import { useLocation, useNavigate, useSearchParams } from "react-router-dom";

interface OptionType {
  value: string;
  label: string;
}
export const NUMPERPAGE = 11
export const CATEGORY_SEARCH_KEY = 'categories'
export const PAGE_SEARCH_KEY = 'page'
export const SEARCH_SEARCH_KEY = 'search'

const Explore: React.FC = () => {
  const [searchParams, setSearchParams] = useSearchParams()
  const [pageNum, setPageNum] = useState(1)
  const [projects, setProjects] = useState<ProjectWId[]>([])
  const [hasNext, setHasHext] = useState<Boolean>(false)
  const categoryChoices = ['business', 'software engineering', 'art'] // TODO get from backend?
  const categorySelectOptions: OptionType[] = categoryChoices.map(opt => ({
    value: opt,
    label: opt
  }))
  const [categories, setCategories] = useState<OptionType[]>([])
  // let categoryValues: OptionType[] = []
  // TODO get state from front page search
  const [searchQuery, setSearchQuery] = useState<string>('')
  
  const detectCategoryChange = (selected: MultiValue<OptionType>) => {
    const selectedValues = selected.map(o => o.value)
    const params = new URLSearchParams(searchParams)
    params.set(CATEGORY_SEARCH_KEY, selectedValues.join(','))
    params.set(PAGE_SEARCH_KEY, '1')
    setSearchParams(params)
  }

  const updatePageNum = (newNum: number) => {
    setPageNum(newNum)
    const params = new URLSearchParams(searchParams)
    params.set(PAGE_SEARCH_KEY, newNum.toString())
    setSearchParams(params)
  }

  useEffect(() => {
    const page = Number(searchParams.get(PAGE_SEARCH_KEY)) || 1
    setPageNum(page)
    let categtmp = searchParams.get(CATEGORY_SEARCH_KEY)
    let categories: string[]
    if (categtmp) {
      categories = categtmp.split(',')
    } else {
      categories = []
    }
    const catOptions: OptionType[] = categories.map(opt => ({
      value: opt,
      label: opt
    }))
    setCategories(catOptions)
    const searchQuery = searchParams.get(SEARCH_SEARCH_KEY) || ''
    setSearchQuery(searchQuery)
    const filters: Filters = { categories, searchQuery }
    axiosBase.post(`/projects?page=${page}&size=${NUMPERPAGE}`, filters)
    .then(res => {
      setProjects(res.data.items || [])
      setHasHext(res.data.hasNext)
    })
    .catch(err => {
      console.log(err)
    })
  }, [searchParams])

  return (
    <div className='explore-parent'>
      <div className='explore-wrapper'>
        <div className='filter-parent'>
          <SearchBar 
            searchParams={searchParams} 
            setSearchParams={setSearchParams}
            searchQuery={searchQuery}
            setSearchQuery={setSearchQuery}
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
                value={categories}
                />
            </div>
          </div>
        </div>
        {projects.length === 0 &&
          <div className='no-projects'> 
            <div>
            no projects found through those filters (yet) 
            </div>
          </div>
        }
        <div className='projectContainer'>
          <CreateProjectTile/>
          {projects.length > 0 && projects.map((value) => (
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
          {
            pageNum > 1 ? 
            <div onClick={()=>updatePageNum(pageNum-1)} 
            className='navButton' id='prev-button'>
                <>
                  <LeftSVG className="SVGarrow"/>
                </> 
            </div>:
            <div className='arrow-placeholder'></div>
          }
          <span className='page-num'>page {pageNum} </span>
          {
            hasNext ?
            <div onClick={()=>updatePageNum(pageNum+1)} 
            className='navButton' id='next-button'>
                <>
                  <RightSVG className="SVGarrow"/>
                </>
            </div> :
            <div className='arrow-placeholder'></div>
          }
        </div>
      </div>
    </div>
  )
}

export default Explore

