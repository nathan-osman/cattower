import BigButton from '../components/BigButton'
import LampSvg from '../images/lamp.svg'

export default function Index() {

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
    <>
      <p>Main controls page</p>
      <br />
      <div className="flex gap-x-4">
        <BigButton
          icon={LampSvg}
          text="Turn LED on"
          onClick={() => sendCommand('top_on')}
        />
        <BigButton
          icon={LampSvg}
          text="Turn LED off"
          onClick={() => sendCommand('top_off')}
        />
      </div>
    </>
  )
}
