import { ButtonHTMLAttributes, PropsWithChildren, useState } from 'react'
import { clsx } from 'clsx'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  className?: string
}

export default function Button(props: PropsWithChildren<Props>) {

  const { className, ...buttonProps } = props

  const [pressed, setPressed] = useState(false)

  const combinedClassName = clsx(
    className,
    "p-2 border rounded shadow-md bg-white",
    {
      'bg-background-selected': pressed,
    },
  )

  return (
    <button
      {...buttonProps}
      onMouseDown={() => setPressed(true)}
      onMouseUp={() => setPressed(false)}
      className={combinedClassName}
    >
      {props.children}
    </button>
  )
}
