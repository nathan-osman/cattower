import { ButtonHTMLAttributes } from 'react'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  text: string
  primary?: boolean
}

export default function Button(props: Props) {

  const { text, ...buttonProps } = props

  return (
    <button
      {...buttonProps}
      className="p-2 border rounded shadow"
    >
      {text}
    </button>
  )
}
