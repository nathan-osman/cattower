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
    "p-2 border rounded shadow-md bg-white flex justify-center gap-x-2",
    {
      '!bg-background-selected': pressed,
    },
  )

  return (
    <button
      {...buttonProps}
      onTouchStart={() => setPressed(true)}
      onTouchEnd={() => setPressed(false)}
      className={combinedClassName}
    >
      {props.children}
    </button>
  )
}
