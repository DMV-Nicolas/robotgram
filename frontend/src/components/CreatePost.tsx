import { useNavigate } from 'react-router-dom'
import { Close } from './Icons'
import './CreatePost.css'
import { useCreatePost } from '../hooks/useCreatePost'
import { useState } from 'react'

export function CreatePost() {
  const { createPost } = useCreatePost()
  const [inputAmount, setInputAmount] = useState([0])
  const navigate = useNavigate()

  const handleGoBack = () => {
    navigate(-1)
  }

  const handleClick = () => {
    setInputAmount((prevInputAmount) => [...prevInputAmount, prevInputAmount[prevInputAmount.length - 1] + 1])
  }

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    const form = e.target as HTMLFormElement
    const formData = new FormData(form)

    const description = formData.get('description') as string
    const images = inputAmount.map((v) => {
      const imageUrl = formData.get(`imageUrl${v}`) as string
      return imageUrl
    })

    createPost({ description, images })
  }

  return (
    <div className="createPostContainer">
      <div className="createPost">
        <h2 className='createPost__title'>Create a new post!</h2>
        <form className='createPost__form' onSubmit={handleSubmit}>
          <ul className='createPost__ul'>
            {inputAmount.map((v) => (
              <li className='createPost__li' key={v}>
                <input className='createPost__input' type="text" name={`imageUrl${v}`} />
              </li>
            ))}
          </ul>
          <textarea className='createPost__textarea' name="description" cols={30} rows={10}></textarea>
          <input className='createPost__input--increment' type="button" name='increment' value="+" onClick={handleClick} />
          <button className='createPost__submit'>submit</button>
        </form>
      </div>
      <button className='postModalContainer__close' onClick={handleGoBack}>
        <Close />
      </button>
    </div>
  )
}
