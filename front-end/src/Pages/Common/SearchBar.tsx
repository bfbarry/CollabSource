import { Dispatch, FC, KeyboardEvent, SetStateAction } from "react";
import './SearchBar.css'
import { PAGE_SEARCH_KEY, SEARCH_SEARCH_KEY } from "../Explore/Explore";
import { SetURLSearchParams, useNavigate } from "react-router-dom";

interface Props {
  searchParams: URLSearchParams
  setSearchParams: SetURLSearchParams
  searchQuery: string
  setSearchQuery: Dispatch<SetStateAction<string>>
  redirect: Boolean
  // onSearch: () => void;
}
const SearchBar: FC<Props> = ({searchParams, setSearchParams, searchQuery, setSearchQuery, redirect}) => {
  const navigate = useNavigate()

  const onSearchEnter = (e: KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      if (redirect) {
        const params = new URLSearchParams(searchParams)
        params.set(SEARCH_SEARCH_KEY, searchQuery)
        params.set(PAGE_SEARCH_KEY, '1')
        navigate(`/explore/?${params.toString()}`)
      } else {
        const params = new URLSearchParams(searchParams)
        params.set(SEARCH_SEARCH_KEY, searchQuery)
        params.set(PAGE_SEARCH_KEY, '1')
        setSearchParams(params)
      }
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