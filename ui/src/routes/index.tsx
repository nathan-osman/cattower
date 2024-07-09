import Button from '../components/Button'

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
      <p>Home page</p>
      <br />
      <Button
        text="Turn LED on"
        onClick={() => sendCommand('top_on')}
      />
      <br />
      <Button
        text="Turn LED off"
        onClick={() => sendCommand('top_off')}
      />
    </>
  )
}
