package ami

import (
	"syscall"
	"unsafe"
	"errors"
	"fmt"
	"strconv"
)

var (
	amiDLL           = syscall.NewLazyDLL("ami32.dll")
	procAMI_REQUEST  = amiDLL.NewProc("AMI_REQUEST")
)

func GetLinkageStatus(status int) (output string) {
	switch status {
		case 00000: output = "สำเร็จ"; break;
		case 90001: output = "ไม่ได้ Login เข้าใช้งานระบบ"; break;
		case 90005: output = "ไม่มีสิทธิในการทำงาน (Time to work Out)"; break;
		case 90007: output = "ไม่มีสิทธิในการทำงาน (Invalid Secret Code)"; break;
		case 90008: output = "ใช้สิทธิในการตรวจสอบข้อมูลครบแล้ว (Quota out of limit)"; break;
		case 90009: output = "ไม่มีสิทธิในการทำงาน (Invalid Smart Card - not found in card_ctl)"; break;
		case 90011: output = "ข้อมูลที่ส่งมาตรวจสอบไม่ถูกต้อง"; break;
		case 90012: output = "ใช้เวลาในการให้บริการมากกว่าที่กำหนดไว้ (Service time out)"; break;
		case 90013: output = "ยังไม่เปิดให้บริการหน่วยงานที่ร้องขอ"; break;
		case 90014: output = "อยู่ระหว่างปรับปรุง"; break;
		case 90015: output = "ไม่เปิดให้บริการ Linkage Center"; break;
		case 90016: output = "Citizen not login"; break;
		case 90090: output = "PIN ไม่ถูกต้อง"; break;
		case 90020: output = "ไม่สามารถทำการค้นหาข้อมูลได้ (Key as Space)"; break;
		case 90025: output = "ไม่สามารถเพิ่มจำนวนการใช้งานได้ (Update Account Error)"; break;
		case 90026: output = "ไม่สามารถติดต่อฐานข้อมูลตรวจสอบสิทธิได้ (Connect DB Check X error)"; break;
		case 90027: output = "ไม่สามารถจัดเก็บรหัสตรวจสอบข้อมูลได้ (Update X error)"; break;
		case 90028: output = "ไม่สามารถตรวจสอบฐานข้อมูลบัตรของผู้ใช้งานได้ (Select iknoemp_card error)"; break;
		case 90029: output = "ไม่สามารถจัดเก็บรหัสตรวจสอบข้อมูลได้ (Update X space error)"; break;
		case 90040: output = "ไม่มีสิทธิในการทำงาน (card_st error)"; break;
		case 90043: output = "ไม่มีสิทธิในการทำงาน (Using Code not match in emp_card)"; break;
		case 90044: output = "ไม่มีสิทธิในการทำงาน (Check SAS error)"; break;
		case 90045: output = "ไม่มีสิทธิในการทำงาน (SAS error - not match)"; break;
		case 90046: output = "ไม่มีสิทธิในการทำงาน (Using Code sened not match in emp_card)"; break;
		case 90050: output = "ไม่มีสิทธิในการทำงาน (Quota as zero)"; break;
		case 90500: output = "ไม่พบข้อมูล"; break;
		case 91001: output = "ไม่สามารถหาข้อมูล EDUCATION ได้"; break;
		case 91002: output = "ไม่สามารถหาข้อมูล SOLDIRE ได้"; break;
		case 95001: output = "ยังไม่ได้ลงทะเบียนผู้ใช้งาน"; break;
		case 95002: output = "บัตรที่เข้าใช้งาน ไม่ใช่บัตรใบล่าสุด"; break;
		case 95003: output = "ค่า Y ที่ส่งมาตรวจสอบไม่ถูกต้อง"; break;
		case 95004: output = "ไม่มีสิทธิในการใช้งานระบบ Linkage Center"; break;
		case 95005: output = "ไม่ระบุรหัสหน่วยงานเพื่อเข้าใช้งาน Linkage Center"; break;
		case 95010: output = "หน่วยงานไม่มีสิทธิในการร้องขอข้อมูลไปยัง Service ปลายทาง"; break;
		case 95011: output = "ไม่มีสิทธิในการร้องขอข้อมูลไปยัง Service ปลายทาง"; break;
		case 95012: output = "ไม่พบ Service ที่ร้องขอข้อมูลในระบบ Linkage Center"; break;
		case 95013: output = "ไม่สามารถร้องขอข้อมูลจาก Linkage Center ไปยัง Service ปลายทางได้"; break;
		case 95014: output = "ไม่มีสิทธิในการร้องขอข้อมูลไปยัง Service ปลายทางด้วยสิทธิประชาชน"; break;
		case 95015: output = "ไม่สามารถส่งข้อมูลได้เนื่องจากมีขนาดเกิน 64 KB"; break;
		case 95016: output = "Service ไม่รองรับการค้นหาข้อมูลด้วย 5000"; break;
		case 95017: output = "Service ไม่รองรับการค้นหาข้อมูลด้วย 6000"; break;
		case 95018: output = "ข้อมูลที่ต้องการค้นหาไม่อยู่ในรูปแบบ URL Encode"; break;
		case 99301: output = "ไม่สามารถตรวจสอบภาพใบหน้าคนต่างด้าวได้"; break;
		case 99304: output = "ไม่พบภาพใบหน้าคนต่างด้าวในฐานข้อมูล"; break;
		case 99305: output = "ไม่สามารถอ่านไฟล์ภาพใบหน้าคนต่างด้าวได้"; break;
		case 99701: output = "ไม่พบรายการการเปลี่ยนแปลงที่อยู่ในฐานข้อมูล"; break;
		case 99702: output = "ไม่พบรายการการเปลี่ยนแปลงชื่อในฐานข้อมูล"; break;
		case 99703: output = "ไม่พบรายการการเปลี่ยนแปลงสัญชาติในฐานข้อมูล"; break;
		case 99706: output = "ไม่สามารถเพิ่มจำนวนการใช้งานได้ (Update Account Error)"; break;
		case 99707: output = "ไม่พบรายการชื่อภาษาอังกฤษในฐานข้อมูล"; break;
		case 99801: output = "ไม่ระบุค่าของบ้านเลขที่และรหัสจังหวัด อำเภอตำบล"; break;
		case 99983: output = "ไม่สามารถส่งภาพใบหน้าได้เนื่องจากมีขนาดเกิน 20 KB"; break;
		case 99989: output = "ไม่สามารถตรวจสอบรายการบัตรได้"; break;
		case 99990: output = "ตรวจสอบข้อมูลแถบแม่เหล็ก เลขควบคุม 2 ไม่ถูกต้อง"; break;
		case 99991: output = "ไม่พบรายการบัตรก่อนหน้า/ถัดไปในฐานข้อมูล"; break;
		case 99992: output = "ไม่พบรายการบัตรในฐานข้อมูล"; break;
		case 99993: output = "ไม่สามารถอ่านไฟล์ภาพใบหน้าได้"; break;
		case 99994: output = "ไม่พบภาพใบหน้าในฐานข้อมูล"; break;
		case 99995: output = "ไม่สามารถตรวจสอบภาพใบหน้าได้"; break;
		case 99997: output = "ไม่สามารถตรวจสอบเลขควบคุมบัตรประจำตัวประชาชนได้"; break;
		case 99999: output = "ไม่มีสิทธิในการทำงาน (Invalid Application ID)"; break;
		default: output = ""; break;
	}
		
	// 00XXX	HTTP Status Code จากบริการข้อมูลของหน่วยงานเชื่อมโยง (ดูความหมายของ HTTP Status Code)

    return
}

func AMI_REQUEST(requestData []byte, replyData []byte) (err error) {
	var status int32 = 0
	var timeout int32 = 1000

	var hostname []byte = []byte("amihost")
	var requestLen int = len(requestData)
	var replyMaxSize int = len(replyData)
	var replyLen int

	ret, _, errno := procAMI_REQUEST.Call(
		uintptr(unsafe.Pointer(&hostname[0])), 
		uintptr(unsafe.Pointer(&requestData[0])), 
		uintptr(unsafe.Pointer(&requestLen)), 
		uintptr(unsafe.Pointer(&replyData[0])), 
		uintptr(unsafe.Pointer(&replyMaxSize)), 
		uintptr(unsafe.Pointer(&replyLen)), 
		uintptr(unsafe.Pointer(&timeout)), 
		uintptr(unsafe.Pointer(&status)) )
	_ = errno
	if ret != 0 {
		e := fmt.Sprintf("[AMI_REQUEST: ret=%v]", ret)
		err = errors.New(e)
		return
	} else {
		err = nil
	}
	return
}

func AMI_REQUEST_9080(pid string, cid string, userOfficeCode string) (randomToken string, err error) {
	var req_9080 []byte = make([]byte, 4+13+16+5)
	code := "9080"
	// Code 4 bytes
	copy(req_9080[0:4], []byte(code)[0:4] )
	// Pid 13 bytes
	copy(req_9080[4:4+13], []byte(pid)[0:13] )
	// Cid 16 bytes
	copy(req_9080[4+13:4+13+16], []byte(cid)[0:16] )
	// office code 5 bytes
	copy(req_9080[4+13+16:], []byte(userOfficeCode)[0:5] )
	
	var reply_9080 []byte = make([]byte, 4+5+32)
	err = AMI_REQUEST(req_9080, reply_9080)
	if err != nil {
		fmt.Printf("AMI_REQUEST_9080 error: %v\r\n", err.Error())
		randomToken = ""
		return
	}

	// r_code := string(reply_9080[0:4])
	r_status := string(reply_9080[4:4+5])
	r_random := string(reply_9080[4+5:4+5+32])

	if r_status != "00000" {
		
		s, err2 := strconv.Atoi(r_status)
		if err2 != nil {
			s = 99999
		}
		e := fmt.Sprintf("[AMI_REQUEST_9080: status=%v, %v]", r_status, GetLinkageStatus(s))
		err = errors.New(e)
		fmt.Printf(err.Error())
		randomToken = ""
		return
	} else {
		randomToken = r_random
	}

	// fmt.Printf("req_9080: %x\n", req_9080)
	// fmt.Printf("reply_9080: %x\n", reply_9080)
	// fmt.Printf("reply_9080 code: %v\n", r_code)
	// fmt.Printf("reply_9080 status: %v\n", r_status)
	// fmt.Printf("reply_9080 random: %v\n", r_random)
	return
}

func AMI_REQUEST_9081(pid string, cid string, randomToken string, envelope string) (token string, err error) {
	var req_9081 []byte = make([]byte, 4+13+16+32+255)
	code := "9081"
	// Code 4 bytes
	copy(req_9081[0:4], []byte(code)[0:4] )
	// Pid 13 bytes
	copy(req_9081[4:4+13], []byte(pid)[0:13] )
	// Cid 16 bytes
	copy(req_9081[4+13:4+13+16], []byte(cid)[0:16] )
	// randomToken 32 bytes
	copy(req_9081[4+13+16:4+13+16+32], []byte(randomToken)[0:32] )
	// envelope 255 bytes
	if len(envelope) > 255 {
		e := fmt.Sprintf("[AMI_REQUEST_9081: envelope_size > maxsize]")
		err = errors.New(e)
		return
	}
	copy(req_9081[4+13+16+32:], []byte(envelope))
	
	var reply_9081 []byte = make([]byte, 4+5+32+32)
	err = AMI_REQUEST(req_9081, reply_9081)
	if err != nil {
		fmt.Printf("AMI_REQUEST_9081 error: %v\r\n", err.Error())
		token = ""
		return
	}

	// r_code := string(reply_9081[0:4])
	r_status := string(reply_9081[4:4+5])
	r_random := string(reply_9081[4+5:4+5+32])
	// r_session := string(reply_9081[4+5+32:4+5+32+32])

	if r_status != "00000" {
		
		s, err2 := strconv.Atoi(r_status)
		if err2 != nil {
			s = 99999
		}
		e := fmt.Sprintf("[AMI_REQUEST_9081: status=%v, %v]", r_status, GetLinkageStatus(s))
		err = errors.New(e)
		fmt.Printf(err.Error())
		token = ""
		return
	} else {
		token = r_random
	}

	// fmt.Printf("req_9081: %x\n", req_9081)
	// fmt.Printf("reply_9081: %x\n", reply_9081)
	// fmt.Printf("reply_9081 status: %v\n", r_status)
	// fmt.Printf("reply_9081 random: %v\n", r_random)
	// fmt.Printf("reply_9081 session: %v\n", r_session)

	return
}

func AMI_REQUEST_5000(token string, officecode string, versioncode string, service string, query_pid string) (output string, err error) {
	var req_5000 []byte = make([]byte, 4+32+5+2+3+13)
	code := "5000"
	// Code 4 bytes
	copy(req_5000[0:4], []byte(code)[0:4] )
	// Token 32 bytes
	copy(req_5000[4:4+32], []byte(token)[0:32] )
	// OfficeCode 5 bytes
	copy(req_5000[4+32:4+32+5], []byte(officecode)[0:5] )
	// versioncode 2 bytes
	copy(req_5000[4+32+5:4+32+5+2], []byte(versioncode)[0:2] )
	// service 3 bytes
	copy(req_5000[4+32+5+2:4+32+5+2+3], []byte(service)[0:3])
	// Pid 13 bytes
	copy(req_5000[4+32+5+2+3:], []byte(query_pid)[0:13])

	var reply_5000 []byte = make([]byte, 4+5+32760)
	err = AMI_REQUEST(req_5000, reply_5000)
	if err != nil {
		fmt.Printf("AMI_REQUEST_5000 error: %v\r\n", err.Error())
		output = ""
		return
	}

	r_status := string(reply_5000[4:4+5])
	r_data := string(reply_5000[4+5:])

	if r_status != "00000" {
		
		s, err2 := strconv.Atoi(r_status)
		if err2 != nil {
			s = 99999
		}
		e := fmt.Sprintf("[AMI_REQUEST_5000: status=%v, %v]", r_status, GetLinkageStatus(s))
		err = errors.New(e)
		fmt.Printf(err.Error())
		output = ""
		return
	} else {
		output = r_data
	}

	return
}

func GetIdData(token string, query_pid string) (output string, err error) {
	officecode := "00023"
	versioncode := "01"
	servicecode := "001"
	return AMI_REQUEST_5000(token, officecode, versioncode , servicecode, query_pid)
}

func GetHomeData(token string, query_pid string) (output string, err error) {
	officecode := "00023"
	versioncode := "01"
	servicecode := "027"
	return AMI_REQUEST_5000(token, officecode, versioncode , servicecode, query_pid)
}
