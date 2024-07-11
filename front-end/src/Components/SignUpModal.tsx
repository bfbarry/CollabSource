import React, { ChangeEvent, Dispatch, SetStateAction, useState } from 'react';
import axios from 'axios';


interface Props {
    setShowSignUp: Dispatch<SetStateAction<Boolean>>
}

interface UserRequestBody {
    name: string;
    password: string;
    email: string;
    description: string;
}

const SignUpModal: React.FC<Props> = ({setShowSignUp}) => { 
    const [ err, setErr ] = useState<string>("");
    const [ formData, setFormData ] = useState<UserRequestBody>({
        name: '',
        password: '',
        email: '',
        description: '',
    })

    const register = async () => {

        try {
        const response = await axios.post(`http://nodeproxy:8000/auth/register`, 
        { email: formData.email, password : formData.password, description: formData.description, name: formData.name });
        } catch (error){
            console.log(error)
            setErr("placeholder error")
        }
    } 

    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData({           
            ...formData,
            [name]: value,
        });
    };

    const modalStyle: React.CSSProperties = {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '5px',
        maxWidth: '80%',
        maxHeight: '80%',
        overflow: 'auto',
        display: 'block'
      };

      const modalOverlayStyle: React.CSSProperties = {
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'
      };

    return(
        <div style={modalOverlayStyle}>
       <div style={modalStyle}>
            {err && 
            <div> {err} </div>
            }
            <button onClick={() => setShowSignUp(false)}>close</button>
            <label>Name (First, Last)</label>
            <input type="text" name="name" value={formData.name} onChange={handleChange}/>
            <br/>
            <label>email</label>
            <input type="text"  name="email" value={formData.email} onChange={handleChange}/>

            <label>description</label>
            <input type="text" name="description" value={formData.description} onChange={handleChange}/>
            <br/>
            <label>password</label>
            <input type="text" name="password" value={formData.password} onChange={handleChange}/>
            <br/>
            <button type="submit" onClick={register}>submit</button>


       </div>
       </div>
    );
}

export default SignUpModal;