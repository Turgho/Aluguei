export const validateCPF = (cpf: string): boolean => {
  // Remove non-numeric characters
  const cleanCPF = cpf.replace(/\D/g, '');
  
  // Check if CPF has 11 digits
  if (cleanCPF.length !== 11) return false;
  
  // Check for known invalid CPFs (all same digits)
  if (/^(\d)\1{10}$/.test(cleanCPF)) return false;
  
  // Calculate first check digit
  let sum = 0;
  for (let i = 0; i < 9; i++) {
    sum += parseInt(cleanCPF[i]) * (10 - i);
  }
  let remainder = sum % 11;
  const firstCheckDigit = remainder < 2 ? 0 : 11 - remainder;
  
  if (firstCheckDigit !== parseInt(cleanCPF[9])) return false;
  
  // Calculate second check digit
  sum = 0;
  for (let i = 0; i < 10; i++) {
    sum += parseInt(cleanCPF[i]) * (11 - i);
  }
  remainder = sum % 11;
  const secondCheckDigit = remainder < 2 ? 0 : 11 - remainder;
  
  return secondCheckDigit === parseInt(cleanCPF[10]);
};