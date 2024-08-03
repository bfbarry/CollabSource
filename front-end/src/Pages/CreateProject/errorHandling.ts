import { FormFieldsError, ProjectForm } from "./CreateProject"
import { Dispatch, SetStateAction } from 'react';

const checkFormError = (noErrorObj: FormFieldsError, formData: ProjectForm, setFormFieldError: Dispatch<SetStateAction<FormFieldsError>>) => {
  let newState: FormFieldsError = { ...noErrorObj }
  const {name, description, tags} = formData
  // TODO define these in some config
  const nameMin=7, descriptionMin=7, tagsArrMin=3, seekingArrMin=1
  const nameMax=100, descriptionMax=1000, tagMax=30, tagsArrMax=10, seekingMax=30, seekingArrMax=50

  if (name === '') {
    newState.nameErr = 'Name cannot be empty'
  } else if (name.length < nameMin) {
    newState.nameErr = `Name must be at least ${nameMin} characters`
  } else if (name.length > nameMax) {
    newState.nameErr = `Name cannot exceed ${nameMax} characters`
  }

  if (description === '') {
    newState.descriptionErr = 'description cannot be empty'
  } else if (description.length < descriptionMin) {
    newState.descriptionErr = `description must be at least ${descriptionMin} characters`
  } else if (description.length > descriptionMax) {
    newState.descriptionErr = `description cannot exceed ${descriptionMax} characters`
  }

  const tagArr = tags.split(',').map(e => e.trim())
  if (tagArr.length < tagsArrMin) {
    newState.tagsErr = `must have at least ${tagsArrMin} tags`
  } else if (tagArr.length > tagsArrMax) {
    newState.tagsErr = `can't have more than ${tagsArrMax} tags`
  }
  for (const e of tagArr) {
    if (e.length > tagMax) {
      newState.tagsErr = `a tag can't have more than ${tagMax} characters`
    }
  }

  const seekingArr = tags.split(',').map(e => e.trim()) // split on any empty string gives ['']
  if (seekingArr.length === 1 && seekingArr[0] === '') {
    newState.seekingErr = `must have at least ${seekingArrMin} role`
  } else if (seekingArr.length > seekingArrMax) {
    newState.seekingErr = `can't have more than ${seekingArrMax} role`
  }
  for (const e of seekingArr) {
    if (e.length > seekingMax) {
      newState.seekingErr = `a role can't have more than ${seekingMax} characters`
    }
  }
  setFormFieldError(newState)
  let errorHappened = false
  Object.entries(newState).forEach(([k, v]) => {
    if (v !== "") {
      errorHappened = true
    }
  })

  return errorHappened
}

export default checkFormError