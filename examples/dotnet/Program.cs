using System.Collections.Generic;
using System.Linq;
using Pulumi;
using baremetal = Pulumi.baremetal;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new baremetal.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

