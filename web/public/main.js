$(document).ready(function () {
    setTimeout(function () {getReport();}, 0);
    setInterval(function () {getReport();}, 1000)
})

// Nothing to see here.
function getReport() {
    $.ajax({
        type: 'GET',
        url: 'https://api.melody.systems/api/v0.1/getreport',
        success: function(data) {
            console.log(data);
            document.getElementById("hostname").innerHTML = data.Host.Hostname;

            document.getElementById("kernelRelease").innerHTML = data.Host.Kernel.Release;

            document.getElementById("kernelVersion").innerHTML = data.Host.Kernel.Version;

            document.getElementById("cpuUtilization").innerHTML = data.CPU.Utilization;

            document.getElementById("cpuLoadAvgOneMin").innerHTML = data.CPU.LoadAvg.OneMin;

            document.getElementById("cpuLoadAvgFiveMin").innerHTML = data.CPU.LoadAvg.FiveMin;

            document.getElementById("cpuLoadAvgFifteenMin").innerHTML = data.CPU.LoadAvg.FifteenMin;
                                
            document.getElementById("memoryTotal").innerHTML = data.Memory.Total;
                                
            document.getElementById("memoryAvailable").innerHTML = data.Memory.Available;
                                
            document.getElementById("memoryPercentUsed").innerHTML = data.Memory.PercentUsed;

            document.getElementById("memorySwapTotal").innerHTML = data.Memory.SwapTotal;
                                
            document.getElementById("memorySwapFree").innerHTML = data.Memory.SwapFree;
                                
            document.getElementById("memorySwapPercentUsed").innerHTML = data.Memory.SwapPercentUsed;

            document.getElementById("networkPublicIP").innerHTML = data.Network.PublicIP;
        }
    })
}