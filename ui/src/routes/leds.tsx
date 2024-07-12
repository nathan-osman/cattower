import BigButton from '../components/BigButton'
import LampSvg from '../images/lamp.svg'

export default function Leds() {

  function sendCommand(cmd: string) {
    fetch('/api/set-colors', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify({
        command: cmd,
      }),
    })
  }

  return (
    <div className="flex gap-x-2">
      <BigButton
        icon={LampSvg}
        text="Top LEDs on"
        onClick={() => sendCommand('top_on')}
      />
      <BigButton
        icon={LampSvg}
        text="Top LEDs off"
        onClick={() => sendCommand('top_off')}
      />
      <BigButton
        icon={LampSvg}
        text="Side LEDs on"
        onClick={() => sendCommand('sides_on')}
      />
      <BigButton
        icon={LampSvg}
        text="Side LEDs off"
        onClick={() => sendCommand('sides_off')}
      />
    </div>
  )
}
