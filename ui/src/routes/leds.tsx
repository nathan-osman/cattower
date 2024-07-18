import { useState } from 'react'
import BigButton from '../components/BigButton'
import ColorChooser from '../components/ColorChooser'
import LampSvg from '../images/lamp.svg'

export default function Leds() {

  const [topColor, setTopColor] = useState('#ffffff')
  const [sidesColor, setSidesColor] = useState('#ffffff')

  function sendCommand(command: string, color?: string) {
    fetch('/api/leds/set-colors', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify({ command, color }),
    })
  }

  return (
    <div className="inline-grid grid-rows-3 grid-flow-col gap-2">
      <BigButton
        icon={LampSvg}
        text="Top LEDs on"
        onClick={() => sendCommand('top_on', topColor)}
      />
      <BigButton
        icon={LampSvg}
        text="Top LEDs off"
        onClick={() => sendCommand('top_off')}
      />
      <ColorChooser
        color={topColor}
        onChange={(c: string) => {
          setTopColor(c)
          sendCommand('top_on', topColor)
        }}
      />
      <BigButton
        icon={LampSvg}
        text="Side LEDs on"
        onClick={() => sendCommand('sides_on', sidesColor)}
      />
      <BigButton
        icon={LampSvg}
        text="Side LEDs off"
        onClick={() => sendCommand('sides_off')}
      />
      <ColorChooser
        color={sidesColor}
        onChange={(c: string) => {
          setSidesColor(c)
          sendCommand('sides_on', sidesColor)
        }}
      />
    </div>
  )
}
