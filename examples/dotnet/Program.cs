using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Baremetal = UnMango.Baremetal;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Baremetal.Random("myRandomResource", new()
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

