import { useState } from 'react'
import './Slider.css'

interface Props {
  id: string
  username: string
  images: string[]
}

export function Slider({ id, username, images }: Props) {
  const [slide, setSlide] = useState(0)

  const prevSlide = (): void => {
    if (slide > 0) setSlide(slide - 1)
  }

  const nextSlide = (): void => {
    if (slide < images.length - 1) setSlide(slide + 1)
  }

  return (
    <div className="slider">
      {slide > 0 &&
        <span className="slider__leftArrow instagramIcons" onClick={prevSlide}></span>
      }
      <img className="slider__image" src={images[slide]} alt={`Post image of ${username}`} />
      {slide < images.length - 1 &&
        <span className="slider__rightArrow instagramIcons" onClick={nextSlide}></span>
      }
      <div className="slider__indicators">
        {
          images.map((_, idx) => (
            <span
              key={`${id}-${idx}`}
              className={`slider__indicator ${slide === idx ? 'slider__indicator--selected' : ''}`}
              onClick={() => { setSlide(idx) }}></span>
          ))
        }
      </div>
    </div>
  )
}
