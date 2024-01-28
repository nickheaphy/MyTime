function validateFormData(formdataobject) {
    // check to make sure all the required fields have been set
    const requiredData = ['start','end','bgcolour','description','customer','primaryLogType','secondaryLogType']
    for (i in requiredData) {
        if (formdataobject.hasOwnProperty(requiredData[i]) != true) return false

    }
    return true
}

// change the colours
function LightenDarkenColor(col, amt) {
    col = parseInt(col, 16);
    return (((col & 0x0000FF) + amt) | ((((col >> 8) & 0x00FF) + amt) << 8) | (((col >> 16) + amt) << 16)).toString(16);
}