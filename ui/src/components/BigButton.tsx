import Button from './Button'

type Props = {
  icon: string
  text: string
  onClick: () => void
}

export default function BigButton(props: Props) {
  return (
    <Button
      type="button"
      className="inline-flex flex-col items-center"
      onClick={props.onClick}
    >
      <img src={props.icon} className="w-24" />
      {props.text}
    </Button>
  )
}
