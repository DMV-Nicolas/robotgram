import { createRef, useEffect, useState } from 'react'
import './Slider.css'

interface Props {
  id: string
  username: string
  images: string[]
}

export function Slider({ id, username, images }: Props) {
  const [slide, setSlide] = useState(0)
  const [sliderHeight, setSliderHeight] = useState(0)
  const sliderRef = createRef<HTMLDivElement>()

  const prevSlide = (): void => {
    if (slide > 0) setSlide(slide - 1)
  }

  const nextSlide = (): void => {
    if (slide < images.length - 1) setSlide(slide + 1)
  }

  useEffect(() => {
    if (sliderRef.current instanceof HTMLDivElement) {
      const image = new Image()
      image.src = images[0]
      const sliderWidth = sliderRef.current.clientWidth

      image.onload = () => {
        const height = image.height / (image.width / sliderWidth)
        setSliderHeight(height)
      }
    }
  }, [])

  return (
    <div className="slider" ref={sliderRef}>
      {slide > 0 &&
        <span className="slider__leftArrow instagramIcons" onClick={prevSlide}></span>
      }
      <img
        style={{ height: `${sliderHeight}px` }}
        className="slider__image"
        src={images[slide]}
        alt={`Post image of ${username}`}
      />
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
