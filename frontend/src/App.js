import './App.css';
import {useEffect, useState} from 'react'

export default function TestPage() {
    const [code, setCode] = useState('')
    const [goRunStdout, setGoRunStdout] = useState('')
    const [historyArr, setHistoryArr] = useState([])


    useEffect(() => {
        setCode(window.localStorage.getItem("code"))
    }, [])


    function doSubmit() {
        console.log(code)
        // 发送 POST 请求
        fetch('/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json', // 设置请求头为 JSON 格式
            },
            body: JSON.stringify({code: code}),
        }).then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`)
            }
            return response.json()
        }).then(data => {
            console.log('Response:', data)
            setGoRunStdout(data.result)
            setHistoryArr([...historyArr, {code: data.code, msg: new Date(Date.now()).toLocaleTimeString()}])
        }).catch(error => {
            console.error('Error:', error)
            setGoRunStdout('Error:' + error)
            setHistoryArr([...historyArr, {code: "error", msg: error.toString()}])
        })
    }

    function doSave() {
        console.log(code)
        window.localStorage.setItem("code", code)
    }

    function doClear() {
        setCode("")
    }


    return (
        <div>
            <h1>GO Editor:</h1>
            <h2>code:</h2>
            <textarea name="" id="" cols="80" rows="10" value={code} onChange={e => setCode(e.target.value)}
                      style={{border: '2px solid red'}}/>

            <br/>

            <button style={{border: '2px solid red'}} onClick={doSubmit}>submit</button>
            <button style={{border: '2px solid red'}} onClick={doSave}>save</button>
            <button style={{border: '2px solid red'}} onClick={doClear}>clear</button>
            <br/>
            <h2>output:</h2>
            <pre style={{border: '2px solid blue', height: 100}} dangerouslySetInnerHTML={{__html: goRunStdout}}/>

            <ul>
                {historyArr.map((v, idx) => (
                    <li key={idx} style={{color: v.code==='fail' ? 'red' : 'blue'}}>
                        {v.code+`\u00A0\u00A0\u00A0\u00A0\u00A0`+v.msg}
                    </li>
                ))}
            </ul>
        </div>
    )
}
