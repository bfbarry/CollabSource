import { Dispatch, FC, KeyboardEvent, SetStateAction } from "react";
import './SearchBar.css'

import { Filters, ProjectWId } from "../../types/project";
import { NUMPERPAGE } from "../Explore/Explore";
import axiosBase from "../../config/axiosConfig";
import { useNavigate } from "react-router-dom";

interface Props {
  searchQuery: string
  setSearchQuery: Dispatch<SetStateAction<string>>
  setProjects: Dispatch<SetStateAction<ProjectWId[]>>
  setHasHext: Dispatch<SetStateAction<Boolean>>
  categories: String[]
  redirect: Boolean
  // onSearch: () => void;
}
const SearchBar: FC<Props> = ({searchQuery, setSearchQuery, categories, setProjects, setHasHext, redirect}) => {
  const navigate = useNavigate()
  const onSearchEnter = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      //TODO navigate to explore 
      const filters: Filters = { categories, searchQuery }
      axiosBase.post(`/projects?page=$1&size=${NUMPERPAGE}`, filters)
        .then(res => {
          const projects = res.data.items || []
          const hasNext = res.data.hasNext
          if (redirect) {
            //TODO
            navigate('/explore', {state: {projects, hasNext, searchQuery}})
          } else {
            setProjects(projects)
            setHasHext(hasNext)
          }
        })
        .catch(err => {
          console.log(err)
        })
      }
  }

  return (
    <label className='search-label'>
      <input 
      className="search-bar" 
      type="text" 
      value={searchQuery}
      onChange={e => setSearchQuery(e.target.value)}
      onKeyDown={onSearchEnter}
      placeholder="Search for projects"/>
    </label>
    )
}

export default SearchBar