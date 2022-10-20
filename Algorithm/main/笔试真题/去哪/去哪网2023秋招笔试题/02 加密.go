/**
加密
密文 = 明文的X次方 mod N
明文 = 密文的D次放 mod N
给出密文和D N,求解明文
 */

package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 解密
 * @param encryptedNumber int整型 待解密数字
 * @param decryption int整型 私钥参数D
 * @param number int整型 私钥参数N
 * @return int整型
 */

// 20分的100%
func Decrypt( encryptedNumber int ,  decryption int ,  number int ) int {
	// write code here
	res := 1
	for i:=0;i<decryption;i++ {
		res = (res* encryptedNumber )% number
	}
	return res
}