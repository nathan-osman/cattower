import { ButtonHTMLAttributes, PropsWithChildren } from 'react'
import { clsx } from 'clsx'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  className?: string
}

export default function Button(props: PropsWithChildren<Props>) {

  const { className, ...buttonProps } = props

  const combinedClassName = clsx(
    className,
    "p-2 border rounded shadow-md bg-white",
  )

  return (
    <button
      {...buttonProps}
      className={combinedClassName}
    >
      {props.children}
    </button>
  )
}
