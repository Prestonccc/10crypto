import {useState, useEffect} from 'react'

interface IState {
    crypto: {
      id: number
      code: string
      name: string
      price: number
    }[]
  }




const useFetch = (url) =>{
    const [data, setData] = useState<IState["crypto"]>([])
    const [isFetching, setFetch] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
      const abortCont = new AbortController();
      async function fetchData() {
        await fetch(url, { signal: abortCont.signal})
        .then(res => res.json())
        .then(data => {
          setData(data);
          setError(null);
          setFetch(false);
        })
        .catch(err => {
          if (err.name === 'AbortError'){
            console.log('fetch aborted')
          } else {
            setFetch(false)
            setError(err.message);
          }
        })        
      }
      fetchData()

      return () => abortCont.abort();
    },[url])

    return {data,error,isFetching}
}

export default useFetch;