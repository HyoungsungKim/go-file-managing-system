async function submit() {
    let userId = document.getElementById("userID").value;
    let title = document.getElementById("title").value;
    let description = document.getElementById("description").value;

    console.log(userId, title, description)

    const object = {
        userId: userId,
        title: title,
        description: description,
    }

    console.log(object)

    const response = await fetch("submit", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(object),
    })

    const serverResponse = await response;
    //console.log("Response: " + JSON.stringify(serverResponse));
    console.log(serverResponse.ok);
    console.log(serverResponse.status);
    console.log(serverResponse);
}