# 4g-lte-sdk
4g-lte-sdk

### doc
```cshapp
 int nPortNo = int.Parse(comPort.Replace("COM", ""));
 int  ret = xxtSms.OpenModem(nPortNo, nBaudRate, "");
 if (ret == 1){
       modemOpen = true;
     // MessageBox.Show("短信猫打开成功");
 }

public void CloseComm(){
    try{
        xxtSms.CloseModem();
        modemOpen = false;
    }
    catch (System.Exception ex){
        //MessageBox.Show(ex.Message);
    }
}

public bool SendMsg(string phone, string msg){
    bool result = false;
    int code = 8;
    if (xxtSms.SendSms(code, phone, msg) == 1)
    {
        result = true;              
        // MessageBox.Show("发送成功");
    }
    return result;

}
```

```

using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;

namespace ATIAN.MapPointsConnectServiceUI.ShortMessage
{
    public class xxtSms
    {
        private const string dllPath = "xxtSmsDll.dll";

        [DllImport(dllPath)]
        public static extern void AutoDelMsg(int val);
        //初始化
        [DllImport(dllPath)]
        public static extern int OpenModem(int nPortNo, int nBaudRate, string sN);

        //发送
        [DllImport(dllPath)]
        public static extern int SendSms(int nBit, string sNo, string sCon);
        //目标手机号码+序列号+已发送次数   例如：  138123456789,001,0

        //接收
        [DllImport(dllPath)]
        public static extern int ReadNewSms(StringBuilder sNo, StringBuilder sCon, StringBuilder sRecvTime);

        //接收
        [DllImport(dllPath)]//, CharSet=CharSet.Ansi)]
        public static extern int ReadSms(int sId, StringBuilder sNo, StringBuilder sCon, StringBuilder sRecvTime);


        //关闭
        [DllImport(dllPath)]
        public static extern int CloseModem();

        [DllImport("kernel32.dll")]
        public static extern int GetTickCount();
    }
}

```