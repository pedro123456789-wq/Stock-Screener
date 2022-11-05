<script>
</script>


<template>
    <div>
        <h1 class = "main-text">Stock Screener</h1>
        
        <div class = "buttons-container">    
            <button class = "option-button">
                Today's Picks
            </button>

            <button class = "option-button">
                Custom Filters 
            </button>
        </div>


        <img class = "main-image" src="../../assets/LandingImage.jpg" />
    </div>
</template>


<style>
    .main-text{
        color: white;
        background: transparent;
        text-align: center;
    }

    .secondary-text{
        color: white;
        background: transparent;
        text-align: center;
    }
    .main-image{
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .buttons-container{
        display: flex;
        position: absolute;
        float: right;
        left: 50%;
        top: 50%;
        transform: translate(-50%,-50%);
        text-align: center;
    }

    .option-button{
        background: black;
        color: white;
        width: 15rem;
        height: 5rem;
        border-radius: 0.5rem;
        margin: 5rem;
        font-size: 1rem;
    }

    .option-button:hover{
        cursor: pointer;
        background: #05fa63;
    }
</style>
