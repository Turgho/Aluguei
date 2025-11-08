import { useState } from 'react';
import { router } from 'expo-router';
import { apiService } from '../../../services/api';
import { useAuth } from '../../../contexts/AuthContext';
import { BRAZILIAN_STATES } from '../../../lib/constants/states';

interface PropertyFormData {
  title: string;
  description: string;
  address: string;
  city: string;
  state: string;
  zipCode: string;
  bedrooms: string;
  bathrooms: string;
  area: string;
  rentAmount: string;
}

export const usePropertyForm = (onStatusChange: (status: { type: 'success' | 'error' | null; message: string }) => void) => {
  const { owner } = useAuth();
  const [form, setForm] = useState<PropertyFormData>({
    title: '',
    description: '',
    address: '',
    city: '',
    state: '',
    zipCode: '',
    bedrooms: '',
    bathrooms: '',
    area: '',
    rentAmount: ''
  });
  const [loading, setLoading] = useState(false);
  const [loadingCep, setLoadingCep] = useState(false);

  const fetchAddressByCep = async (cep: string) => {
    const cleanCep = cep.replace(/\D/g, '');
    if (cleanCep.length !== 8) return;

    setLoadingCep(true);
    try {
      const response = await fetch(`https://viacep.com.br/ws/${cleanCep}/json/`);
      const data = await response.json();
      
      if (!data.erro) {
        setForm(prev => ({
          ...prev,
          address: data.logradouro || prev.address,
          city: data.localidade || prev.city,
          state: data.uf || prev.state
        }));
      }
    } catch (error) {
      console.log('Erro ao buscar CEP:', error);
    } finally {
      setLoadingCep(false);
    }
  };

  const updateForm = (field: keyof PropertyFormData, value: string) => {
    if (field === 'state') {
      const upperValue = value.toUpperCase();
      if (upperValue.length <= 2 && (upperValue === '' || BRAZILIAN_STATES.some(state => state.startsWith(upperValue)))) {
        setForm(prev => ({ ...prev, [field]: upperValue }));
      }
    } else {
      setForm(prev => ({ ...prev, [field]: value }));
      
      if (field === 'zipCode') {
        const cleanCep = value.replace(/\D/g, '');
        if (cleanCep.length === 8) {
          fetchAddressByCep(cleanCep);
        }
      }
    }
  };

  const handleSubmit = async () => {
    if (!form.title || !form.address || !form.city || !form.state || !form.rentAmount) {
      onStatusChange({ type: 'error', message: 'Por favor, preencha todos os campos obrigatórios' });
      return;
    }

    if (!BRAZILIAN_STATES.includes(form.state as any)) {
      onStatusChange({ type: 'error', message: 'Estado inválido. Use apenas siglas dos estados brasileiros (ex: SP, RJ)' });
      return;
    }

    if (!owner?.id) {
      onStatusChange({ type: 'error', message: 'Usuário não autenticado' });
      return;
    }

    setLoading(true);
    onStatusChange({ type: null, message: '' });
    
    try {
      // Mock delay para visualizar loading
      await new Promise(resolve => setTimeout(resolve, 2000));
      await apiService.createProperty({
        owner_id: owner.id,
        title: form.title.trim(),
        description: form.description.trim(),
        address: form.address.trim(),
        city: form.city.trim(),
        state: form.state.trim().toUpperCase(),
        zip_code: form.zipCode.replace(/\D/g, ''),
        bedrooms: form.bedrooms ? parseInt(form.bedrooms) : 0,
        bathrooms: form.bathrooms ? parseInt(form.bathrooms) : 0,
        area: form.area ? parseInt(form.area) : 0,
        rent_amount: parseFloat(form.rentAmount.replace(/[^\d,]/g, '').replace(',', '.')),
        status: 'available'
      });
      
      onStatusChange({ type: 'success', message: 'Propriedade criada com sucesso!' });
      setTimeout(() => router.replace('./dashboard'), 2000);
    } catch (error: any) {
      onStatusChange({ type: 'error', message: error.message || 'Erro ao criar propriedade' });
    } finally {
      setLoading(false);
    }
  };

  return {
    form,
    updateForm,
    loading,
    loadingCep,
    handleSubmit
  };
};