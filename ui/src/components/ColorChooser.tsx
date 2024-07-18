import { useState } from 'react'
import Button from './Button'
import { clsx } from 'clsx'

type Props = {
  color: string
  onChange: (color: string) => void
}

const colors = [
  '#ffffff',
  '#ff0000',
  '#00ff00',
  '#0000ff',
  '#ffff00',
  '#00ffff',
  '#ff00ff',
]

export default function ColorChooser(props: Props) {

  const [showPopup, setShowPopup] = useState(false)

  const className = clsx(
    'absolute',
    'border border-foreground bg-background',
    'grid grid-rows-2 grid-flow-col',
    'gap-4 p-4 z-10',
    {
      'invisible': !showPopup,
    },
  )

  function handleClick(c: string) {
    setShowPopup(false)
    props.onChange(c)
  }

  return (
    <div className="relative">
      <div className={className}>
        {
          colors.map(v => (
            <div
              className="w-8 h-8 border border-black"
              style={{ background: v }}
              onClick={() => handleClick(v)}
            />
          ))
        }
      </div>
      <Button
        className="flex items-center w-full"
        onClick={() => setShowPopup(true)}
      >
        <div
          className="w-4 h-4 border border-black"
          style={{ background: props.color }}
        />
        Color...
      </Button>
    </div>
  )
}
