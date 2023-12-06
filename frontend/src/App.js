import './App.css'
import React, {useEffect, useState} from 'react'
import {Button, Input, Layout, Space, Timeline} from 'antd'
import {Content, Header} from 'antd/es/layout/layout'

const {TextArea} = Input
export default function TestPage() {
    const [code, setCode] = useState('')
    const [goRunStdout, setGoRunStdout] = useState('')
    const [historyArr, setHistoryArr] = useState([])

    useEffect(() => {
        setCode(window.localStorage.getItem('code'))
    }, [])

    function doSubmit() {
        console.log(code)
        fetch('http://localhost:8080/golang/run-code', {
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
            setGoRunStdout(data.data)
            setHistoryArr([
                {
                    children: <Space size={50}>
                        <div>{new Date(Date.now()).toLocaleTimeString()}</div>
                        <div style={{fontSize: 16}}>{data.message}</div>
                    </Space>,
                    color: data.message === 'fail' ? 'red' : 'green'
                },
                ...historyArr
            ])
        }).catch(error => {
            console.error('Error:', error)
            setGoRunStdout('Error:' + error)
            setHistoryArr([...historyArr, {code: 'error', msg: error.toString()}])
        })
    }

    function doSave() {
        console.log(code)
        window.localStorage.setItem('code', code)
    }

    function doClear() {
        setCode('')
    }

    return (
        <Layout>
            <Header style={{
                background: '#ffffff',
                boxShadow: '0 1px 2px 0 rgba(0, 0, 0, 0.03), 0 1px 6px -1px rgba(0, 0, 0, 0.02), 0 2px 4px 0 rgba(0, 0, 0, 0.02)',
                fontSize: 32,
                fontWeight: 'bold',
                textAlign: 'center'
            }}>
                GO Editor
            </Header>
            <Content style={{padding: 24, minHeight: 'calc(100vh - 64px)'}}>
                <div style={{border: '1px solid #d9d9d9', borderRadius: '12px',}}>
                    <div style={{
                        padding: '8px 12px',
                        borderBlockEnd: '1px solid #d9d9d9',
                        fontSize: 24,
                        fontWeight: 'bold',

                    }}>
                        Code:
                    </div>
                    <TextArea
                        style={{background: 'transparent'}}
                        bordered={false}
                        rows={10}
                        value={code}
                        onChange={e => setCode(e.target.value)}
                    />
                </div>
                <Space size={16} style={{marginTop: 16}}>
                    <Button type="primary" onClick={doSubmit}>submit</Button>
                    <Button type="primary" onClick={doSave}>save</Button>
                    <Button type="primary" onClick={doClear}>clear</Button>
                </Space>
                <div style={{border: '1px solid #d9d9d9', borderRadius: '12px', marginTop: 24}}>
                    <div style={{
                        padding: '8px 12px',
                        borderBlockEnd: '1px solid #d9d9d9',
                        fontSize: 24,
                        fontWeight: 'bold',
                    }}>
                        Output:
                    </div>
                    <TextArea
                        style={{background: 'transparent',color:'#000000'}}
                        rows={5}
                        disabled={true}
                        bordered={false}
                        value={goRunStdout}
                    />
                </div>
                <div style={{border: '1px solid #d9d9d9', borderRadius: '12px', marginTop: 24}}>
                    <div style={{
                        padding: '8px 12px',
                        borderBlockEnd: '1px solid #d9d9d9',
                        fontSize: 12,
                        fontWeight: 'bold',
                    }}>
                        Execution records:
                    </div>
                    <div style={{height: 200,padding:24, overflowY: 'auto',}}>
                        <Timeline items={historyArr?.slice(0, 5) ?? []}/>
                    </div>

                </div>
            </Content>
        </Layout>
    )
}
