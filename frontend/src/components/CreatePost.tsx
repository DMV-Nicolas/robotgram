import { useNavigate } from 'react-router-dom'
import { Close, Url } from './Icons'
import { useCreatePost } from '../hooks/useCreatePost'
import { useId, useState } from 'react'
import './CreatePost.css'
import { toast } from 'sonner'
import { validateUrl } from '../services/validate'

export function CreatePost() {
  const { createPost } = useCreatePost()
  const [inputAmount, setInputAmount] = useState([0])
  const inputImageUrlID = useId()
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

    let error = false
    const images = inputAmount.map((v) => {
      const imageUrl = formData.get(`imageUrl${v}`) as string
      if (!validateUrl(imageUrl)) {
        toast.error('Invalid Url!')
        error = true
      }
      return imageUrl
    })

    if (error) return
    createPost({ description, images })
  }

  return (
    <div className="createPostContainer">
      <div className="createPost">
        <h2 className='createPost__title'>Create a new post!</h2>
        <form className='createPost__form' onSubmit={handleSubmit}>
          <textarea className='createPost__input createPost__input--textarea' name="description" cols={30} rows={4} placeholder='Description'></textarea>
          <ul className='createPost__ul'>
            {inputAmount.map((v) => (
              <li className='createPost__inputField' key={v}>
                <label className="createPost__label" htmlFor={inputImageUrlID + v}>
                  <Url />
                </label>
                {inputAmount.length === v + 1
                  ? <>
                    <input className='createPost__input' type="text" id={inputImageUrlID + v} name={`imageUrl${v}`} placeholder={`Image url ${v + 1}`} />
                  </>
                  : <>
                    <input className='createPost__input createPost__input--shorter' type="text" id={inputImageUrlID + v} name={`imageUrl${v}`} placeholder={`Image url ${v + 1}`} />
                    <input className='createPost__input createPost__input--increment' type="button" name='increment' value="+" onClick={handleClick} />
                  </>
                }
              </li>
            ))}
          </ul>
          <button className='createPost__submit'>submit</button>
        </form>
      </div>
      <button className='postModalContainer__close' onClick={handleGoBack}>
        <Close />
      </button>
    </div>
  )
}
